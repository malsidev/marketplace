from sqlalchemy import select
from models.models import Categories
from sqlalchemy.ext.asyncio import AsyncSession

async def addCategories(data, db: AsyncSession):
    categories = Categories(
        name = data.name
    )
    db.add(categories)

    await db.commit()
    await db.refresh(categories)

    return {
        "id": categories.id
    }

async def categories(db: AsyncSession):
    stmt = (select(Categories))

    res = await db.execute(stmt)
    data = res.scalars().all()


    return [
        {
            "id": categori.id,
            "name" : categori.name
        }
        for categori in data
    ]
    