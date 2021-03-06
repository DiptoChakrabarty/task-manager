package database

import (
	"fmt"
	//"github.com/DiptoChakrabart/task-manager/utils"
	bolt "go.etcd.io/bbolt"
)

var bucketlist = []byte("bucketlist")

func CreateBucket(bucketName []byte) error {
	err := CreateBucketListBucket()
	if err != nil {
		return err
	}
	err = AddBucketToBucketList(bucketName)
	if err != nil {
		return err
	}
	return DbConn.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			return err
		}
		//fmt.Println("Bucket Done")
		return err
	})
}

func DeleteABucket(bucketName []byte) error {
	err := DeleteBucketFromBucketList(bucketName)

	if err != nil {
		return err
	}

	return DbConn.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket(bucketName)
		if err != nil {
			return err
		}
		return nil
	})
}

func CreateBucketListBucket() error {
	return DbConn.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketlist)
		if err != nil {
			return err
		}
		//fmt.Println("BucketList made")
		return nil
	})
}

func AddBucketToBucketList(bucketName []byte) error {
	//fmt.Println("Created BucketList")
	//fmt.Println(string(bucketName))
	//fmt.Println(string(bucketlist))
	return DbConn.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(bucketlist)
		return bkt.Put(bucketName, []byte{1})
	})
}

func DeleteBucketFromBucketList(bucketName []byte) error {
	err := DbConn.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(bucketlist)
		return bkt.Delete(bucketName)
	})
	if err != nil {
		return err
	}
	return nil
}

func GetAllBuckets() ([]string, error) {
	var buckets []string
	err := DbConn.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(bucketlist)
		cur := bkt.Cursor()

		for k, _ := cur.First(); k != nil; k, _ = cur.Next() {
			fmt.Println(string(k))
			buckets = append(buckets, string(k))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return buckets, nil
}
