from fastapi import APIRouter, Header

from services.products import products
from schemas.schemas import Products

router = APIRouter(prefix="/products", tags=["products"])

@router.post("")
async def post_products(
        product_data : Products,
        user_id: str | None = Header(default=None, alias="X-User-ID")
        
):
    return await products(user_id)