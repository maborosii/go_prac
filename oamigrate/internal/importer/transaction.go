package importer

import "xorm.io/xorm"

type dbTransactionFunc func(sess *xorm.Session) error

func InTransaction(engine *xorm.Engine, callback dbTransactionFunc) error {
	var err error

	sess := engine.NewSession()
	defer sess.Close()

	if err = sess.Begin(); err != nil {
		return err
	}

	//将sess传入回调函数执行sql操作，此回调函数将返回err，如果sql操作有错则err不为空，将会在commit之前实现回滚
	err = callback(sess)
	if err != nil {
		sess.Rollback()
		return err
	} else if err = sess.Commit(); err != nil {
		return err
	}

	return nil
}
