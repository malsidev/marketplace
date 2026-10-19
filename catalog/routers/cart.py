from fastapi import APIRouter
from schemas.schemas import CartProductRequest

router = APIRouter(prefix="/cart", tags=["cart"])

@router.post("")
async def get_data_cart(data: CartProductRequest):
    return {"mess" : data}