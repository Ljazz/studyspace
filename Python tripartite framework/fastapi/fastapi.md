# FastAPI基础

FastAPI是一个用于构建API的现代、快速（高性能）的web框架

关键特性
- 快速
- 高效编码：提高功能开发速度
- 更少bug：减少认为导致错误
- 智能：极佳的编译器支持
- 简单：设计的易于使用和学习，阅读文档的时间更短
- 简短：使代码重复最小化。通过不同的参数神功实现丰富功能。bug更少
- 健壮：生产可用级别的代码。还有自动生成的交互式文档。
- 标准化：基于（并完全兼容）API的相关开放标准OpenAPI和JSON Schema

## 安装FastAPI
---
> pip install fastapi
> 
> pip install uvicorn

---

## 简单示例

```python
from fastapi import FastAPI

app = FastAPI()

@app.get("/")
async def root():
    return {"message": "Hello world"}

@app.get("/items/{item_id}")
def read_item(item_id: int, q: Optional[str] = None):
    return {"item_id": item_id, "q": q}
```

**运行实时服务器**

![运行实时服务器](./images/运行实时服务器.png)

`uvicorn main:app --reload`命令含义如下
- `main`：`main.py`文件
- `app`：在`main.py`文件中通过`app = FastAPI()`
- `--reload`：让服务器在更新代码后重启。

**交互式API文档**

访问`http://127.0.0.1:8000/docs`，可以看到自动生成的交互式API文档（由Swagger UI生成）

**可选的API文档**

访问`http://127.0.0.1:8000/redoc`，可以看到另外一个自动生成的文档（由ReDoc生成）

## 示例升级

```python
from typing import Optional
from fastapi import FastAPI
from pydantic import BaseModel

app = FastAPI()


class Item(BaseModel):
    name: str
    price: float
    is_offer: Optional[bool] = None


@app.get("/")
def read_root():
    return {"Hello": "World"}


@app.get("/items/{item_id}")
def read_item(item_id: int, q: Optional[str]=None):
    return {"item_id": item_id, "q": q}


@app.put("/items/{item_id}")
def update_item(item_id: int, item: Item):
    return {"item_name": item.name, "item_id": item_id}
```

