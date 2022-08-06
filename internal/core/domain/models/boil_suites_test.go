// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Engagements", testEngagements)
}

func TestDelete(t *testing.T) {
	t.Run("Engagements", testEngagementsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Engagements", testEngagementsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Engagements", testEngagementsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Engagements", testEngagementsExists)
}

func TestFind(t *testing.T) {
	t.Run("Engagements", testEngagementsFind)
}

func TestBind(t *testing.T) {
	t.Run("Engagements", testEngagementsBind)
}

func TestOne(t *testing.T) {
	t.Run("Engagements", testEngagementsOne)
}

func TestAll(t *testing.T) {
	t.Run("Engagements", testEngagementsAll)
}

func TestCount(t *testing.T) {
	t.Run("Engagements", testEngagementsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Engagements", testEngagementsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Engagements", testEngagementsInsert)
	t.Run("Engagements", testEngagementsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Engagements", testEngagementsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Engagements", testEngagementsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Engagements", testEngagementsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Engagements", testEngagementsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Engagements", testEngagementsSliceUpdateAll)
}
