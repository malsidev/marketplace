from fastapi import APIRouter

router = APIRouter(prefix="/cart", tags=["cart"])

@router.get("/{product_id}")
async def get_data_cart(product_id: int):
    return {"mess" : product_id}