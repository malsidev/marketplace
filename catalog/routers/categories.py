from fastapi import APIRouter, Header

from services.categories import categories
from schemas.schemas import Categories

router = APIRouter(prefix="/categories", tags=["categories"])

@router.post("")
async def post_categories(
        categories_data : Categories,
        user_id: str | None = Header(default=None, alias="X-User-ID")
        
):
    return await categories(categories_data)