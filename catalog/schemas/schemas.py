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