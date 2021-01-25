Django的ORM字段与数据库实际字段的对一个关系

| ORM字段 | 数据库实际字段 |
| --- | --- |
| AutoField | integer AUTO_INCREMENT |
| BigAutoField | bigint AUTO_INCREMENT |
| BinaryField | longblob |
| BooleanField | bool |
| CharField | varchar(%(max_length)s) |
| CommaSeparatedIntegerField | varchar(%(max_length)s) |
| DateField | date |
| DateTimeField | datetime |
| DecimalField | numeric(%(max_digits)s, %(decimal_places)s) |
| DurationField | bigint |
| FileField | varchar(%(max_length)s) |
| FilePathField | varchar(%(max_length)s) |
| FloatField | double precision |
| IntegerField | integer |
| BigIntegerField | bigint |
| IPAddressField | char(15) |
| GenericIPAddressField | char(39) |
| NullBooleanField | bool |
| OneToOneField | integer |
| PositiveIntegerField | integer UNSIGNED |
| PositiveSmallIntegerField | integer UNSIGNED |
| SlugField | varchar(%(max_length)s) |
| SmallIntegerField | smallint |
| TextField | longtext |
| TimeField | time |
| UUIDField | char(32) |
