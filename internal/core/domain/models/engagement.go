// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Engagement is an object representing the database table.
type Engagement struct {
	ID             int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	Action         string    `boil:"action" json:"action" toml:"action" yaml:"action"`
	ActionTime     time.Time `boil:"action_time" json:"action_time" toml:"action_time" yaml:"action_time"`
	Browser        string    `boil:"browser" json:"browser" toml:"browser" yaml:"browser"`
	BrowserVersion string    `boil:"browser_version" json:"browser_version" toml:"browser_version" yaml:"browser_version"`
	Platform       string    `boil:"platform" json:"platform" toml:"platform" yaml:"platform"`
	Identifier     string    `boil:"identifier" json:"identifier" toml:"identifier" yaml:"identifier"`
	Host           string    `boil:"host" json:"host" toml:"host" yaml:"host"`
	Path           string    `boil:"path" json:"path" toml:"path" yaml:"path"`
	FullURL        string    `boil:"full_url" json:"full_url" toml:"full_url" yaml:"full_url"`
	ViewPort       string    `boil:"view_port" json:"view_port" toml:"view_port" yaml:"view_port"`
	Os             string    `boil:"os" json:"os" toml:"os" yaml:"os"`
	OsVersion      string    `boil:"os_version" json:"os_version" toml:"os_version" yaml:"os_version"`
	CreatedAt      null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt      null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *engagementR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L engagementL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EngagementColumns = struct {
	ID             string
	Action         string
	ActionTime     string
	Browser        string
	BrowserVersion string
	Platform       string
	Identifier     string
	Host           string
	Path           string
	FullURL        string
	ViewPort       string
	Os             string
	OsVersion      string
	CreatedAt      string
	UpdatedAt      string
}{
	ID:             "id",
	Action:         "action",
	ActionTime:     "action_time",
	Browser:        "browser",
	BrowserVersion: "browser_version",
	Platform:       "platform",
	Identifier:     "identifier",
	Host:           "host",
	Path:           "path",
	FullURL:        "full_url",
	ViewPort:       "view_port",
	Os:             "os",
	OsVersion:      "os_version",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

var EngagementTableColumns = struct {
	ID             string
	Action         string
	ActionTime     string
	Browser        string
	BrowserVersion string
	Platform       string
	Identifier     string
	Host           string
	Path           string
	FullURL        string
	ViewPort       string
	Os             string
	OsVersion      string
	CreatedAt      string
	UpdatedAt      string
}{
	ID:             "engagement.id",
	Action:         "engagement.action",
	ActionTime:     "engagement.action_time",
	Browser:        "engagement.browser",
	BrowserVersion: "engagement.browser_version",
	Platform:       "engagement.platform",
	Identifier:     "engagement.identifier",
	Host:           "engagement.host",
	Path:           "engagement.path",
	FullURL:        "engagement.full_url",
	ViewPort:       "engagement.view_port",
	Os:             "engagement.os",
	OsVersion:      "engagement.os_version",
	CreatedAt:      "engagement.created_at",
	UpdatedAt:      "engagement.updated_at",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var EngagementWhere = struct {
	ID             whereHelperint64
	Action         whereHelperstring
	ActionTime     whereHelpertime_Time
	Browser        whereHelperstring
	BrowserVersion whereHelperstring
	Platform       whereHelperstring
	Identifier     whereHelperstring
	Host           whereHelperstring
	Path           whereHelperstring
	FullURL        whereHelperstring
	ViewPort       whereHelperstring
	Os             whereHelperstring
	OsVersion      whereHelperstring
	CreatedAt      whereHelpernull_Time
	UpdatedAt      whereHelpernull_Time
}{
	ID:             whereHelperint64{field: "\"engagement\".\"id\""},
	Action:         whereHelperstring{field: "\"engagement\".\"action\""},
	ActionTime:     whereHelpertime_Time{field: "\"engagement\".\"action_time\""},
	Browser:        whereHelperstring{field: "\"engagement\".\"browser\""},
	BrowserVersion: whereHelperstring{field: "\"engagement\".\"browser_version\""},
	Platform:       whereHelperstring{field: "\"engagement\".\"platform\""},
	Identifier:     whereHelperstring{field: "\"engagement\".\"identifier\""},
	Host:           whereHelperstring{field: "\"engagement\".\"host\""},
	Path:           whereHelperstring{field: "\"engagement\".\"path\""},
	FullURL:        whereHelperstring{field: "\"engagement\".\"full_url\""},
	ViewPort:       whereHelperstring{field: "\"engagement\".\"view_port\""},
	Os:             whereHelperstring{field: "\"engagement\".\"os\""},
	OsVersion:      whereHelperstring{field: "\"engagement\".\"os_version\""},
	CreatedAt:      whereHelpernull_Time{field: "\"engagement\".\"created_at\""},
	UpdatedAt:      whereHelpernull_Time{field: "\"engagement\".\"updated_at\""},
}

// EngagementRels is where relationship names are stored.
var EngagementRels = struct {
}{}

// engagementR is where relationships are stored.
type engagementR struct {
}

// NewStruct creates a new relationship struct
func (*engagementR) NewStruct() *engagementR {
	return &engagementR{}
}

// engagementL is where Load methods for each relationship are stored.
type engagementL struct{}

var (
	engagementAllColumns            = []string{"id", "action", "action_time", "browser", "browser_version", "platform", "identifier", "host", "path", "full_url", "view_port", "os", "os_version", "created_at", "updated_at"}
	engagementColumnsWithoutDefault = []string{"action", "browser", "browser_version", "platform", "identifier", "host", "path", "full_url", "view_port", "os", "os_version", "created_at", "updated_at"}
	engagementColumnsWithDefault    = []string{"id", "action_time"}
	engagementPrimaryKeyColumns     = []string{"id"}
)

type (
	// EngagementSlice is an alias for a slice of pointers to Engagement.
	// This should almost always be used instead of []Engagement.
	EngagementSlice []*Engagement
	// EngagementHook is the signature for custom Engagement hook methods
	EngagementHook func(context.Context, boil.ContextExecutor, *Engagement) error

	engagementQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	engagementType                 = reflect.TypeOf(&Engagement{})
	engagementMapping              = queries.MakeStructMapping(engagementType)
	engagementPrimaryKeyMapping, _ = queries.BindMapping(engagementType, engagementMapping, engagementPrimaryKeyColumns)
	engagementInsertCacheMut       sync.RWMutex
	engagementInsertCache          = make(map[string]insertCache)
	engagementUpdateCacheMut       sync.RWMutex
	engagementUpdateCache          = make(map[string]updateCache)
	engagementUpsertCacheMut       sync.RWMutex
	engagementUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var engagementBeforeInsertHooks []EngagementHook
var engagementBeforeUpdateHooks []EngagementHook
var engagementBeforeDeleteHooks []EngagementHook
var engagementBeforeUpsertHooks []EngagementHook

var engagementAfterInsertHooks []EngagementHook
var engagementAfterSelectHooks []EngagementHook
var engagementAfterUpdateHooks []EngagementHook
var engagementAfterDeleteHooks []EngagementHook
var engagementAfterUpsertHooks []EngagementHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Engagement) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range engagementBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Engagement) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range engagementBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Engagement) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range engagementBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Engagement) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range engagementBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Engagement) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range engagementAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Engagement) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range engagementAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Engagement) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range engagementAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Engagement) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range engagementAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Engagement) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range engagementAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEngagementHook registers your hook function for all future operations.
func AddEngagementHook(hookPoint boil.HookPoint, engagementHook EngagementHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		engagementBeforeInsertHooks = append(engagementBeforeInsertHooks, engagementHook)
	case boil.BeforeUpdateHook:
		engagementBeforeUpdateHooks = append(engagementBeforeUpdateHooks, engagementHook)
	case boil.BeforeDeleteHook:
		engagementBeforeDeleteHooks = append(engagementBeforeDeleteHooks, engagementHook)
	case boil.BeforeUpsertHook:
		engagementBeforeUpsertHooks = append(engagementBeforeUpsertHooks, engagementHook)
	case boil.AfterInsertHook:
		engagementAfterInsertHooks = append(engagementAfterInsertHooks, engagementHook)
	case boil.AfterSelectHook:
		engagementAfterSelectHooks = append(engagementAfterSelectHooks, engagementHook)
	case boil.AfterUpdateHook:
		engagementAfterUpdateHooks = append(engagementAfterUpdateHooks, engagementHook)
	case boil.AfterDeleteHook:
		engagementAfterDeleteHooks = append(engagementAfterDeleteHooks, engagementHook)
	case boil.AfterUpsertHook:
		engagementAfterUpsertHooks = append(engagementAfterUpsertHooks, engagementHook)
	}
}

// One returns a single engagement record from the query.
func (q engagementQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Engagement, error) {
	o := &Engagement{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for engagement")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Engagement records from the query.
func (q engagementQuery) All(ctx context.Context, exec boil.ContextExecutor) (EngagementSlice, error) {
	var o []*Engagement

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Engagement slice")
	}

	if len(engagementAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Engagement records in the query.
func (q engagementQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count engagement rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q engagementQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if engagement exists")
	}

	return count > 0, nil
}

// Engagements retrieves all the records using an executor.
func Engagements(mods ...qm.QueryMod) engagementQuery {
	mods = append(mods, qm.From("\"engagement\""))
	return engagementQuery{NewQuery(mods...)}
}

// FindEngagement retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEngagement(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Engagement, error) {
	engagementObj := &Engagement{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"engagement\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, engagementObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from engagement")
	}

	if err = engagementObj.doAfterSelectHooks(ctx, exec); err != nil {
		return engagementObj, err
	}

	return engagementObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Engagement) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no engagement provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(engagementColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	engagementInsertCacheMut.RLock()
	cache, cached := engagementInsertCache[key]
	engagementInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			engagementAllColumns,
			engagementColumnsWithDefault,
			engagementColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(engagementType, engagementMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(engagementType, engagementMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"engagement\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"engagement\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into engagement")
	}

	if !cached {
		engagementInsertCacheMut.Lock()
		engagementInsertCache[key] = cache
		engagementInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Engagement.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Engagement) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	engagementUpdateCacheMut.RLock()
	cache, cached := engagementUpdateCache[key]
	engagementUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			engagementAllColumns,
			engagementPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update engagement, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"engagement\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, engagementPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(engagementType, engagementMapping, append(wl, engagementPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update engagement row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for engagement")
	}

	if !cached {
		engagementUpdateCacheMut.Lock()
		engagementUpdateCache[key] = cache
		engagementUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q engagementQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for engagement")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for engagement")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EngagementSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), engagementPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"engagement\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, engagementPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in engagement slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all engagement")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Engagement) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no engagement provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(engagementColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	engagementUpsertCacheMut.RLock()
	cache, cached := engagementUpsertCache[key]
	engagementUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			engagementAllColumns,
			engagementColumnsWithDefault,
			engagementColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			engagementAllColumns,
			engagementPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert engagement, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(engagementPrimaryKeyColumns))
			copy(conflict, engagementPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"engagement\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(engagementType, engagementMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(engagementType, engagementMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert engagement")
	}

	if !cached {
		engagementUpsertCacheMut.Lock()
		engagementUpsertCache[key] = cache
		engagementUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Engagement record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Engagement) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Engagement provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), engagementPrimaryKeyMapping)
	sql := "DELETE FROM \"engagement\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from engagement")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for engagement")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q engagementQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no engagementQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from engagement")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for engagement")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EngagementSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(engagementBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), engagementPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"engagement\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, engagementPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from engagement slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for engagement")
	}

	if len(engagementAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Engagement) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEngagement(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EngagementSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EngagementSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), engagementPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"engagement\".* FROM \"engagement\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, engagementPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in EngagementSlice")
	}

	*o = slice

	return nil
}

// EngagementExists checks if the Engagement row exists.
func EngagementExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"engagement\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if engagement exists")
	}

	return exists, nil
}
