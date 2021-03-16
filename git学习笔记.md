# 1. git是什么？

git是一个分布式版本控制系统。

# 2. git安装

## 2.1 Linux

### Debian/Ubuntu

```bash
apt install git
```

### CentOS

```bash
yum install git
```

### 源码安装

git下载地址：

https://mirrors.edge.kernel.org/pub/software/scm/git/

```bash
wget https://mirrors.edge.kernel.org/pub/software/scm/git/git-2.9.5.tar.gz

tar -zxvf git-2.9.5.tar.gz

cd git-2.9.5

./configure --prefix=/usr/local/git

make && make install
```

## 2.2 Windows安装

安装包下载地址：https://gitforwindows.org/

官网慢，可以用国内的镜像：https://npm.taobao.org/mirrors/git-for-windows/。

## 2.3 Mac安装

Homebrew安装

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

git安装：`brew install git`

# 3. git相关操作

## 3.1 创建仓库命令

### 1. `git init`

【说明】：初始化Git仓库

【格式】：`git init`

【示例】：`git init`

### 2. `git clone`

【说明】：拷贝一个 Git 仓库到本地

【格式】：`git init [url]`

【示例】：`git clone https://github.com/tianqixin/runoob-git-test`

## 3.2 提交与修改

### 1. `git add`

【说明】：添加一个或多个文件（文件夹）到暂存区

【格式】：`git add [file1(dir)] [file2(dir)] ...`

【示例】：`git add 1.txt 2.txt` or `git add dir1 dir2`

【注意】：`git add .`添加当前目录下的所有文件到暂存区

### 2. `git status`

【说明】：查看仓库当前的状态，显示有变更的文件

【格式】：`git status [-s]`

【示例】：`git status`

【注意】：`git status -s`中的`-s`参数可以获得简短的输出结果

### 3. `git diff`

【说明】：比较文件的不同，即暂存区和工作区的差异；显示已写入暂存区和已经被修改但尚未写入暂存区文件对区别

【格式】：`git diff`

【场景】

- 尚未缓存的改动：`git diff`
- 查看已缓存的改动：`git diff --cached`
- 查看已缓存的与未缓存的所有改动：`git diff HEAD`
- 显示摘要而非整个diff：`git diff --stat`

【示例】：

- 显示暂存区和工作区的差异：`git diff [file]`
- 显示暂存区和上一次提交(commit)的差异：`git diff --cached [file]`或`git diff --staged [file]`
- 显示两次提交之间的差异：`git diff [first-branch] ... [second-branch]`

### 4. `git commit`

【说明】：提交暂存区到本地仓库

【格式】：`git commit -m [message`

【示例】

- 提交暂存区到本地仓库：`git commit -m [message]`
  - message是一些备注信息
- 提交缓存区的指定文件到仓库区：`git commit [file1] [file2] ... -m [message]
- `-a`参数设置修改文件后不需要执行`git add`命令，直接来提交：`git commit -a -m [message]`

### 5. `git reset`

【说明】：回退版本，可以指定退回某一次提交的版本。

【格式】：`git reset [--soft | --mixed | --hard] [HEAD]`

- `--mixed`为默认，可以不用带该参数，用于重置暂存区的文件与上一次的提交（commit）保持一致，工作区文件内容保持不变
- `--soft`参数用于回退到某个版本
- `--hard`参数撤销工作区中所有未提交的修改内容，将暂存区与工作区都回到上一次版本，并删除之前的所有信息提交

【示例】：

```bash
git reset HEAD^             # 回退所有内容到上一个版本
git reset HEAD^ hello.py    # 回退hello.py文件的版本到上一个版本
git reset 053e              # 回退到指定的版本

git reset –hard HEAD~3          # 回退上上上一个版本  
git reset –hard bae128          # 回退到某个版本回退点之前的所有信息。 
git reset --hard origin/master  # 将本地的状态回退到和远程的一样 
```

【注意】

- HEAD 说明：
  - HEAD 表示当前版本
  - HEAD^ 上一个版本
  - HEAD^^ 上上一个版本
  - HEAD^^^ 上上上一个版本
  - 以此类推...
- 可以使用 ～数字表示
  - HEAD~0 表示当前版本
  - HEAD~1 上一个版本
  - HEAD^2 上上一个版本
  - HEAD^3 上上上一个版本
  - 以此类推...

### 6. `git rm`

【说明】：删除工作区文件

【格式】：`git rm [file]`

【示例】：

- 将文件从暂存区和工作区中删除
  - 从暂存区和工作区中删除runoob.txt文件：`git rm runoob.txt`
- 如果删除之前修改过并且已经放到暂存区域的话，则必须要用强制删除选项`-f`。
  - 强行从暂存区和工作区中删除修改后的runoob.txt文件：`git rm -f runoob.txt`
- 若想把文件从暂存区移除，但当前工作目录中希望保留，即从跟踪清单中删除，使用`--cached`选项：`git rm --cached [file]`
  - 从暂存区中删除runoob.txt文件：`git rm --cached runoob.txt`
- 可以递归删除：`git rm -r [file/dir]`
  - 若后是目录做参数，则会递归删除整个目录中的所有子目录和文件

### 7. `git mv`

【说明】：用于移动或重命名一个文件、目录或软连接。

【格式】：`git mv [file] [newfile]`

【示例】

- 新文件名已存在，但还是要重命名它，可以使用`-f`参数：`git mv -f [file] [newfile]`

## 3.3 提交日志

### 1. `git log`

【说明】：查看历史提交记录

【格式】：`git log`

【注意】

- 可以使用`--online`选项来查看历史记录的简洁版本：`git log --oneline`
- 可以用 `--graph` 选项，查看历史中什么时候出现了分支、合并：``
- 可以用 `--reverse` 参数来逆向显示所有日志：`git log --reverse --oneline`
- 查找指定用户的提交日志可以使用命令：`git log --author=xxx`
- 指定日期，可以执行几个选项：--since 和 --before，但是也可以用 --until 和 --after。

### 2. `git blame [file]`

【说明】：以列表形式查看指定文件的历史修改记录。

【格式】：`git blame [file]`

【示例】：`git blame README`

## 3.4 远程操作

### 1. `git remote`

### 2. `git fetch`

### 3. `get pull`

### 4. `get push`
