package store

import (
	"encoding/json"
	"example.com/inspection14/internal/domain"
	"fmt"
	"go.etcd.io/bbolt"
	"path/filepath"
	"sync"
)

var buckets = [][]byte{[]byte("vehicles"), []byte("inspections"), []byte("photos"), []byte("reviews"), []byte("archives"), []byte("events")}

type Store struct {
	db *bbolt.DB
	mu sync.RWMutex
}

func Open(path string) (*Store, error) {
	db, err := bbolt.Open(filepath.Clean(path), 0600, nil)
	if err != nil {
		return nil, err
	}
	s := &Store{db: db}
	err = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, e := tx.CreateBucketIfNotExists(b); e != nil {
				return e
			}
		}
		return nil
	})
	if err != nil {
		db.Close()
		return nil, err
	}
	return s, nil
}
func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.db == nil {
		return nil
	}
	e := s.db.Close()
	s.db = nil
	return e
}
func encode(v any) ([]byte, error) { return json.Marshal(v) }
func (s *Store) put(bucket, key string, v any) error {
	data, e := encode(v)
	if e != nil {
		return e
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return fmt.Errorf("store closed")
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte(bucket)).Put([]byte(key), data) })
}
func (s *Store) get(bucket, key string, out any) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return fmt.Errorf("store closed")
	}
	return s.db.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket([]byte(bucket)).Get([]byte(key))
		if v == nil {
			return bbolt.ErrBucketNotFound
		}
		return json.Unmarshal(v, out)
	})
}
func (s *Store) List(bucket string) ([][]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var out [][]byte
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(bucket)).ForEach(func(k, v []byte) error {
			if v != nil {
				out = append(out, append([]byte(nil), v...))
			}
			return nil
		})
	})
	return out, e
}
func (s *Store) SaveVehicle(v domain.Vehicle) error { return s.put("vehicles", v.ID, v) }
func (s *Store) Vehicle(id string) (domain.Vehicle, error) {
	var v domain.Vehicle
	e := s.get("vehicles", id, &v)
	return v, e
}
func (s *Store) SaveInspection(v domain.Inspection) error { return s.put("inspections", v.ID, v) }
func (s *Store) Inspection(id string) (domain.Inspection, error) {
	var v domain.Inspection
	e := s.get("inspections", id, &v)
	return v, e
}
func (s *Store) SavePhoto(v domain.Photo) error           { return s.put("photos", v.ID, v) }
func (s *Store) SaveReview(v domain.Review) error         { return s.put("reviews", v.ID, v) }
func (s *Store) SaveArchive(v domain.ArchiveRecord) error { return s.put("archives", v.ID, v) }
func (s *Store) SaveEvent(v domain.StatusEvent) error     { return s.put("events", v.ID, v) }
