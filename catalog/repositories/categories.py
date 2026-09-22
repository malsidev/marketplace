from models.models import Categories
from sqlalchemy.ext.asyncio import AsyncSession
from fastapi import HTTPException, status

async def addCategories(data, db: AsyncSession):
    brand = Categories(
        name = data.name
    )
    db.add(brand)

    await db.commit()
    await db.refresh(brand)

    return {
        "id": brand.id
    }
