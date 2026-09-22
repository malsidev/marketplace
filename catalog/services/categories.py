from repositories.categories import addCategories


async def categories(data, db):
    return await addCategories(data, db)