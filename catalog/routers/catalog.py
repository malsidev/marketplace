from fastapi import APIRouter, Header

from services.catalog import catalog
from schemas.schemas import Products

router = APIRouter(prefix="/catalog", tags=["catalog"])

@router.get("")
async def get_catalog(
        product_data : Products,
        user_id: str | None = Header(default=None, alias="X-User-ID"),
):
    return await catalog(product_data)
