package mynotes

// pgAdmin4 下操作数据库学习笔记
/*
创建数据库表：
	create table student (id int not null,name varchar(100) not null ,age int not null,sex varchar not null)

删除数据库表：
	drop table 表名

插入数据：
	insert into student (id ,name ,age , sex ) values (1,'Tom',23,'boy')

插入多个数据：
	insert into student values (6,'小明',21,'boy'),(7,'小丽',22,'girl')

查询student表全部数据：
	select * from student

查询student某个列全部数据：
	select 列名 from student
	select id from student

mysql where 语句：
	select id,age from student where name='Tom' 通过名字筛选

更新数据：
update student set id=3,age=18 where name='Tom' 通过名字筛选来修改其他列数据

删除数据表数据：(只是删除数据，没有删除列)
	delete from student  //删除表中所有数据
	delete from student where name='tom' 删除name为tom的所有数据

（alter使用）删除数据库表列：（删除数据删除列）
	alter table student drop column sex;

修改字段长度：
	alter table student modify column name varchar(88);

修改表名称：
	alter table student change colum name lastname varchar(88);

添加表列：
	alter table student add column name varchar(18);

修改表名：
	alter table 原表名 rename to 新表名

mysql Like使用（like语句）：
	select * from student where sex like '%bo'
	输出所有sex列下包含bo字段的数据

mysql 排序（order by 语句）：
	ASC升序，DESC降序
	select * from student order by name ASC

mysql 分组：（group by 语句）
	group by 语句根据一个或者多个列对结果集进行分组
	count()函数，输出某个字段出现次数
	select sex,count(*) from student group by sex 输出sex中各个字段出现次数


把两个表数据连接起来显示：
	select * from 表1 inser join 表2；
	筛选具体数据连接：
	select a.列名称1,a.列名称2,b.列名称1 from 表1 a inser join 表2 ;
	若有相同列：
	select a.列名称1,a.列名称2,b.列名称1 from 表1 a inser join 表2 on a.列名称2=b.列名称2;

NULL处理：
	IS NULL: 当列的值是NULL,此运算符返回true。
	IS NOT NULL: 当列的值不为NULL, 运算符返回true。
	<=>: 比较操作符（不同于=运算符），当比较的的两个值为NULL时返回true。
	select * from student where name is not null;

创建临时表：(退出数据库的时候自动销毁)
	create temporary table table_name;


*/
