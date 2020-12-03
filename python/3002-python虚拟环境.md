# 虚拟环境

## virtualenv

### 安装

---
pip install virtualenv

---

### 创建虚拟环境

---
virtualenv ENV

---

上述`ENV`是要创建的虚拟环境的名称。上述命令会在当前工作目录下，创建一个新的名为`ENV`的目录。

Unix和Linux中，`ENV`目录下包含以下几个目录

- `bin/`：包含新的Python可执行文件和其他提供的脚本/可执行文件。
- `lib/`和`include/`：包含虚拟环境中新Python的支持库文件。新的Python包将会安装在`ENV/lib/pythonX.Y/site-packages/`中

Windows中，`ENV`目录的子目录为`Scripts/`、`Libs/`和`include/`。

### 虚拟环境激活

Unix/Linux中：

---
source ENV/lib/activate

---

window环境

---
ENV/Scripts/activate

---


## venv

可以使用`pyvenv ENV`来创建新的虚拟环境。其余的操作跟virtualenv一致。
