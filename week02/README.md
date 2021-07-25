* Q：我们在数据库操作的时候，比如dao层中当遇到一个sql.ErrNoRows的时候，是否应该Wrap这个error，抛给上层？
* A：应该
