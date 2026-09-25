from repositories.category import addCategories, categories


async def getCategories(db):
    return await categories(db)

async def postCategories(data, db):
    return await addCategories(data, db)