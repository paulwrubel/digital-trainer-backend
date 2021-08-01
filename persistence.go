package main

import "github.com/sirupsen/logrus"

type persistenceObject interface {
	Save(*logrus.Entry, *appData) error
	Get(*logrus.Entry, *appData) error
	Update(*logrus.Entry, *appData) error
	Delete(*logrus.Entry, *appData) error
	Exists(*logrus.Entry, *appData) (bool, error)

	Type() string
}
