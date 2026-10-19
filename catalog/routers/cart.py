from fastapi import APIRouter, Depends
from schemas.schemas import CartProductRequest
from repositories.cart import getCartProduct
from sqlalchemy.ext.asyncio import AsyncSession
from database.pgdb import get_db


router = APIRouter(prefix="/cart", tags=["cart"])

@router.post("")
async def get_data_cart(data: CartProductRequest, db: AsyncSession = Depends(get_db)):
    product = [item.product_id for item in data.items]
    return await getCartProduct(product, db)