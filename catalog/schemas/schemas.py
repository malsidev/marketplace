from pydantic import BaseModel

class Products(BaseModel):
    name : str
    description : str
    brand_id : int
    category_id : int

class Brands(BaseModel):
    name: str

class Categories(BaseModel):
    name: str

class CartProduct(BaseModel):
    product_id : int
    quantity : int

class CartProductRequest(BaseModel):
    items: list[CartProduct]

class ProductID(BaseModel):
    product_id : list[int]