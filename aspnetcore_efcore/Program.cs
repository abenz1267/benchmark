using System.ComponentModel.DataAnnotations.Schema;
using Microsoft.EntityFrameworkCore;

var builder = WebApplication.CreateBuilder(args);
var connectionString = builder.Configuration.GetConnectionString("DefaultConnection");
builder.Services.AddDbContext<Db>(options =>
    options.UseNpgsql(connectionString));

var app = builder.Build();

app.MapGet("/", async (Db db) => await db.Datas.ToListAsync());

app.Run();

class Db: DbContext {
    public Db(DbContextOptions<Db> options): base(options) {}

    public DbSet<Data> Datas => Set<Data>();
}

[Table("posts")]
record Data(int id) {
    public int id {get;set;} = id;
    public string title {get;set;} = default!;
    public string content {get;set;} = default!;
}
