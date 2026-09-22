from fastapi import APIRouter, Header, Depends
from sqlalchemy.ext.asyncio import AsyncSession
from database import get_db
from services.categories import categories
from schemas.schemas import Categories

router = APIRouter(prefix="/categories", tags=["categories"])

@router.post("")
async def post_categories(
        categories_data : Categories,
        db: AsyncSession = Depends(get_db),
        user_id: str | None = Header(default=None, alias="X-User-ID")
        
):
    return await categories(categories_data, db)