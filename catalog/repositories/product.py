from sqlalchemy import select
from models.models import Products
from sqlalchemy.ext.asyncio import AsyncSession

async def addProduct(name, description, brand_id, category_id, db: AsyncSession):
    categories = Products(
        name = name,
        description = description,
        brand_id = brand_id,
        category_id = category_id
    )
    db.add(categories)

    await db.commit()
    await db.refresh(categories)

    return categories.id