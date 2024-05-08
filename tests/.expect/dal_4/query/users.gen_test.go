// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
	"gorm.io/gen/tests/.gen/dal_4/model"
)

func init() {
	InitializeDB()
	err := _gen_test_db.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.User{}) fail: %s", err)
	}
}

func Test_userQuery(t *testing.T) {
	user := newUser(_gen_test_db)
	user = *user.As(user.TableName())
	_do := user.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(user.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <users> fail:", err)
		return
	}

	_, ok := user.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from user success")
	}

	err = _do.Create(&model.User{})
	if err != nil {
		t.Error("create item in table <users> fail:", err)
	}

	err = _do.Save(&model.User{})
	if err != nil {
		t.Error("create item in table <users> fail:", err)
	}

	err = _do.CreateInBatches([]*model.User{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <users> fail:", err)
	}

	_, err = _do.Select(user.ALL).Take()
	if err != nil {
		t.Error("Take() on table <users> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <users> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <users> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <users> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.User{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <users> fail:", err)
	}

	_, err = _do.Select(user.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <users> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <users> fail:", err)
	}

	_, err = _do.Select(user.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <users> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <users> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <users> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <users> fail:", err)
	}

	_, err = _do.ScanByPage(&model.User{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <users> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <users> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <users> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), "id")

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <users> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <users> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <users> fail:", err)
	}
}

var UserFindByUsersTestCase = []TestCase{}

func Test_user_FindByUsers(t *testing.T) {
	user := newUser(_gen_test_db)
	do := user.WithContext(context.Background()).Debug()

	for i, tt := range UserFindByUsersTestCase {
		t.Run("FindByUsers_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.FindByUsers(tt.Input.Args[0].(model.User))
			assert(t, "FindByUsers", res1, tt.Expectation.Ret[0])
		})
	}
}

var UserFindByComplexIfTestCase = []TestCase{}

func Test_user_FindByComplexIf(t *testing.T) {
	user := newUser(_gen_test_db)
	do := user.WithContext(context.Background()).Debug()

	for i, tt := range UserFindByComplexIfTestCase {
		t.Run("FindByComplexIf_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.FindByComplexIf(tt.Input.Args[0].(*model.User))
			assert(t, "FindByComplexIf", res1, tt.Expectation.Ret[0])
		})
	}
}

var UserFindByIfTimeTestCase = []TestCase{}

func Test_user_FindByIfTime(t *testing.T) {
	user := newUser(_gen_test_db)
	do := user.WithContext(context.Background()).Debug()

	for i, tt := range UserFindByIfTimeTestCase {
		t.Run("FindByIfTime_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.FindByIfTime(tt.Input.Args[0].(time.Time))
			assert(t, "FindByIfTime", res1, tt.Expectation.Ret[0])
		})
	}
}

var UserTestForTestCase = []TestCase{}

func Test_user_TestFor(t *testing.T) {
	user := newUser(_gen_test_db)
	do := user.WithContext(context.Background()).Debug()

	for i, tt := range UserTestForTestCase {
		t.Run("TestFor_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.TestFor(tt.Input.Args[0].([]string))
			assert(t, "TestFor", res1, tt.Expectation.Ret[0])
			assert(t, "TestFor", res2, tt.Expectation.Ret[1])
		})
	}
}

var UserTestForKeyTestCase = []TestCase{}

func Test_user_TestForKey(t *testing.T) {
	user := newUser(_gen_test_db)
	do := user.WithContext(context.Background()).Debug()

	for i, tt := range UserTestForKeyTestCase {
		t.Run("TestForKey_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.TestForKey(tt.Input.Args[0].([]string), tt.Input.Args[1].(string), tt.Input.Args[2].(string))
			assert(t, "TestForKey", res1, tt.Expectation.Ret[0])
			assert(t, "TestForKey", res2, tt.Expectation.Ret[1])
		})
	}
}

var UserTestForOrTestCase = []TestCase{}

func Test_user_TestForOr(t *testing.T) {
	user := newUser(_gen_test_db)
	do := user.WithContext(context.Background()).Debug()

	for i, tt := range UserTestForOrTestCase {
		t.Run("TestForOr_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.TestForOr(tt.Input.Args[0].([]string))
			assert(t, "TestForOr", res1, tt.Expectation.Ret[0])
			assert(t, "TestForOr", res2, tt.Expectation.Ret[1])
		})
	}
}

var UserTestIfInForTestCase = []TestCase{}

func Test_user_TestIfInFor(t *testing.T) {
	user := newUser(_gen_test_db)
	do := user.WithContext(context.Background()).Debug()

	for i, tt := range UserTestIfInForTestCase {
		t.Run("TestIfInFor_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.TestIfInFor(tt.Input.Args[0].([]string), tt.Input.Args[1].(string))
			assert(t, "TestIfInFor", res1, tt.Expectation.Ret[0])
			assert(t, "TestIfInFor", res2, tt.Expectation.Ret[1])
		})
	}
}

var UserTestForInIfTestCase = []TestCase{}

func Test_user_TestForInIf(t *testing.T) {
	user := newUser(_gen_test_db)
	do := user.WithContext(context.Background()).Debug()

	for i, tt := range UserTestForInIfTestCase {
		t.Run("TestForInIf_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.TestForInIf(tt.Input.Args[0].([]string), tt.Input.Args[1].(string))
			assert(t, "TestForInIf", res1, tt.Expectation.Ret[0])
			assert(t, "TestForInIf", res2, tt.Expectation.Ret[1])
		})
	}
}

var UserTestForInWhereTestCase = []TestCase{}

func Test_user_TestForInWhere(t *testing.T) {
	user := newUser(_gen_test_db)
	do := user.WithContext(context.Background()).Debug()

	for i, tt := range UserTestForInWhereTestCase {
		t.Run("TestForInWhere_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.TestForInWhere(tt.Input.Args[0].([]string), tt.Input.Args[1].(string), tt.Input.Args[2].(string))
			assert(t, "TestForInWhere", res1, tt.Expectation.Ret[0])
			assert(t, "TestForInWhere", res2, tt.Expectation.Ret[1])
		})
	}
}

var UserTestForUserListTestCase = []TestCase{}

func Test_user_TestForUserList(t *testing.T) {
	user := newUser(_gen_test_db)
	do := user.WithContext(context.Background()).Debug()

	for i, tt := range UserTestForUserListTestCase {
		t.Run("TestForUserList_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.TestForUserList(tt.Input.Args[0].(*[]model.User), tt.Input.Args[1].(string))
			assert(t, "TestForUserList", res1, tt.Expectation.Ret[0])
			assert(t, "TestForUserList", res2, tt.Expectation.Ret[1])
		})
	}
}

var UserTestForMapTestCase = []TestCase{}

func Test_user_TestForMap(t *testing.T) {
	user := newUser(_gen_test_db)
	do := user.WithContext(context.Background()).Debug()

	for i, tt := range UserTestForMapTestCase {
		t.Run("TestForMap_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.TestForMap(tt.Input.Args[0].(map[string]string), tt.Input.Args[1].(string))
			assert(t, "TestForMap", res1, tt.Expectation.Ret[0])
			assert(t, "TestForMap", res2, tt.Expectation.Ret[1])
		})
	}
}

var UserTestIfInIfTestCase = []TestCase{}

func Test_user_TestIfInIf(t *testing.T) {
	user := newUser(_gen_test_db)
	do := user.WithContext(context.Background()).Debug()

	for i, tt := range UserTestIfInIfTestCase {
		t.Run("TestIfInIf_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.TestIfInIf(tt.Input.Args[0].(string))
			assert(t, "TestIfInIf", res1, tt.Expectation.Ret[0])
		})
	}
}

var UserTestMoreForTestCase = []TestCase{}

func Test_user_TestMoreFor(t *testing.T) {
	user := newUser(_gen_test_db)
	do := user.WithContext(context.Background()).Debug()

	for i, tt := range UserTestMoreForTestCase {
		t.Run("TestMoreFor_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.TestMoreFor(tt.Input.Args[0].([]string), tt.Input.Args[1].([]int))
			assert(t, "TestMoreFor", res1, tt.Expectation.Ret[0])
		})
	}
}

var UserTestMoreFor2TestCase = []TestCase{}

func Test_user_TestMoreFor2(t *testing.T) {
	user := newUser(_gen_test_db)
	do := user.WithContext(context.Background()).Debug()

	for i, tt := range UserTestMoreFor2TestCase {
		t.Run("TestMoreFor2_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.TestMoreFor2(tt.Input.Args[0].([]string), tt.Input.Args[1].([]int))
			assert(t, "TestMoreFor2", res1, tt.Expectation.Ret[0])
		})
	}
}

var UserTestForInSetTestCase = []TestCase{}

func Test_user_TestForInSet(t *testing.T) {
	user := newUser(_gen_test_db)
	do := user.WithContext(context.Background()).Debug()

	for i, tt := range UserTestForInSetTestCase {
		t.Run("TestForInSet_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.TestForInSet(tt.Input.Args[0].([]model.User))
			assert(t, "TestForInSet", res1, tt.Expectation.Ret[0])
		})
	}
}

var UserTestInsertMoreInfoTestCase = []TestCase{}

func Test_user_TestInsertMoreInfo(t *testing.T) {
	user := newUser(_gen_test_db)
	do := user.WithContext(context.Background()).Debug()

	for i, tt := range UserTestInsertMoreInfoTestCase {
		t.Run("TestInsertMoreInfo_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.TestInsertMoreInfo(tt.Input.Args[0].([]model.User))
			assert(t, "TestInsertMoreInfo", res1, tt.Expectation.Ret[0])
		})
	}
}

var UserTestIfElseForTestCase = []TestCase{}

func Test_user_TestIfElseFor(t *testing.T) {
	user := newUser(_gen_test_db)
	do := user.WithContext(context.Background()).Debug()

	for i, tt := range UserTestIfElseForTestCase {
		t.Run("TestIfElseFor_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.TestIfElseFor(tt.Input.Args[0].(string), tt.Input.Args[1].([]model.User))
			assert(t, "TestIfElseFor", res1, tt.Expectation.Ret[0])
		})
	}
}

var UserTestForLikeTestCase = []TestCase{}

func Test_user_TestForLike(t *testing.T) {
	user := newUser(_gen_test_db)
	do := user.WithContext(context.Background()).Debug()

	for i, tt := range UserTestForLikeTestCase {
		t.Run("TestForLike_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.TestForLike(tt.Input.Args[0].([]string))
			assert(t, "TestForLike", res1, tt.Expectation.Ret[0])
		})
	}
}

var UserAddUserTestCase = []TestCase{}

func Test_user_AddUser(t *testing.T) {
	user := newUser(_gen_test_db)
	do := user.WithContext(context.Background()).Debug()

	for i, tt := range UserAddUserTestCase {
		t.Run("AddUser_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.AddUser(tt.Input.Args[0].(string), tt.Input.Args[1].(int))
			assert(t, "AddUser", res1, tt.Expectation.Ret[0])
			assert(t, "AddUser", res2, tt.Expectation.Ret[1])
		})
	}
}

var UserAddUser1TestCase = []TestCase{}

func Test_user_AddUser1(t *testing.T) {
	user := newUser(_gen_test_db)
	do := user.WithContext(context.Background()).Debug()

	for i, tt := range UserAddUser1TestCase {
		t.Run("AddUser1_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.AddUser1(tt.Input.Args[0].(string), tt.Input.Args[1].(int))
			assert(t, "AddUser1", res1, tt.Expectation.Ret[0])
			assert(t, "AddUser1", res2, tt.Expectation.Ret[1])
		})
	}
}

var UserAddUser2TestCase = []TestCase{}

func Test_user_AddUser2(t *testing.T) {
	user := newUser(_gen_test_db)
	do := user.WithContext(context.Background()).Debug()

	for i, tt := range UserAddUser2TestCase {
		t.Run("AddUser2_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.AddUser2(tt.Input.Args[0].(string), tt.Input.Args[1].(int))
			assert(t, "AddUser2", res1, tt.Expectation.Ret[0])
		})
	}
}

var UserAddUser3TestCase = []TestCase{}

func Test_user_AddUser3(t *testing.T) {
	user := newUser(_gen_test_db)
	do := user.WithContext(context.Background()).Debug()

	for i, tt := range UserAddUser3TestCase {
		t.Run("AddUser3_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.AddUser3(tt.Input.Args[0].(string), tt.Input.Args[1].(int))
			assert(t, "AddUser3", res1, tt.Expectation.Ret[0])
		})
	}
}

var UserAddUser4TestCase = []TestCase{}

func Test_user_AddUser4(t *testing.T) {
	user := newUser(_gen_test_db)
	do := user.WithContext(context.Background()).Debug()

	for i, tt := range UserAddUser4TestCase {
		t.Run("AddUser4_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.AddUser4(tt.Input.Args[0].(string), tt.Input.Args[1].(int))
			assert(t, "AddUser4", res1, tt.Expectation.Ret[0])
		})
	}
}

var UserAddUser5TestCase = []TestCase{}

func Test_user_AddUser5(t *testing.T) {
	user := newUser(_gen_test_db)
	do := user.WithContext(context.Background()).Debug()

	for i, tt := range UserAddUser5TestCase {
		t.Run("AddUser5_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.AddUser5(tt.Input.Args[0].(string), tt.Input.Args[1].(int))
			assert(t, "AddUser5", res1, tt.Expectation.Ret[0])
		})
	}
}

var UserAddUser6TestCase = []TestCase{}

func Test_user_AddUser6(t *testing.T) {
	user := newUser(_gen_test_db)
	do := user.WithContext(context.Background()).Debug()

	for i, tt := range UserAddUser6TestCase {
		t.Run("AddUser6_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.AddUser6(tt.Input.Args[0].(string), tt.Input.Args[1].(int))
			assert(t, "AddUser6", res1, tt.Expectation.Ret[0])
			assert(t, "AddUser6", res2, tt.Expectation.Ret[1])
		})
	}
}

var UserFindByIDTestCase = []TestCase{}

func Test_user_FindByID(t *testing.T) {
	user := newUser(_gen_test_db)
	do := user.WithContext(context.Background()).Debug()

	for i, tt := range UserFindByIDTestCase {
		t.Run("FindByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.FindByID(tt.Input.Args[0].(int))
			assert(t, "FindByID", res1, tt.Expectation.Ret[0])
		})
	}
}

var UserLikeSearchTestCase = []TestCase{}

func Test_user_LikeSearch(t *testing.T) {
	user := newUser(_gen_test_db)
	do := user.WithContext(context.Background()).Debug()

	for i, tt := range UserLikeSearchTestCase {
		t.Run("LikeSearch_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.LikeSearch(tt.Input.Args[0].(string))
			assert(t, "LikeSearch", res1, tt.Expectation.Ret[0])
		})
	}
}

var UserInSearchTestCase = []TestCase{}

func Test_user_InSearch(t *testing.T) {
	user := newUser(_gen_test_db)
	do := user.WithContext(context.Background()).Debug()

	for i, tt := range UserInSearchTestCase {
		t.Run("InSearch_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.InSearch(tt.Input.Args[0].([]string))
			assert(t, "InSearch", res1, tt.Expectation.Ret[0])
		})
	}
}

var UserColumnSearchTestCase = []TestCase{}

func Test_user_ColumnSearch(t *testing.T) {
	user := newUser(_gen_test_db)
	do := user.WithContext(context.Background()).Debug()

	for i, tt := range UserColumnSearchTestCase {
		t.Run("ColumnSearch_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.ColumnSearch(tt.Input.Args[0].(string), tt.Input.Args[1].([]string))
			assert(t, "ColumnSearch", res1, tt.Expectation.Ret[0])
		})
	}
}
