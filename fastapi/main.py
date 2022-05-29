from starlette.responses import JSONResponse
from fastapi import FastAPI
import asyncpg
import os

app = FastAPI()

connection_pool = None


@app.on_event("startup")
async def setup_database():
    global connection_pool
    connection_pool = await asyncpg.create_pool(
        user=os.getenv('PGUSER', 'postgres'),
        password=os.getenv('PGPASS', 'postgres'),
        database='benchmark',
        host='localhost',
        port=5432
    )

READ = 'SELECT "id","title","content" FROM "posts" WHERE id=$1'
WRITE = 'UPDATE "posts" SET title=$1 WHERE id=$2'


@app.get('/update/20')
async def update():
    worlds = [{}] * 21

    async with connection_pool.acquire() as connection:
        statement = await connection.prepare(READ)
        for i in range(1, 21, 1):
            worlds[i] = dict(await statement.fetchrow(i))
            worlds[i]["title"] = "testtitle"
            await connection.execute(WRITE, "testtitle", i)

    return JSONResponse(worlds)
