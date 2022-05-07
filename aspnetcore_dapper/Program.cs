using System.Data;
using Dapper.Contrib.Extensions;

var builder = WebApplication.CreateBuilder(args);
builder.Services.AddScoped<GetConnection>(
        sp =>
    async () =>
    {
        string connectionString = sp.GetService<IConfiguration>()["DefaultConnection"];
        var connection = new Npgsql.NpgsqlConnection(connectionString);
        await connection.OpenAsync();
        return connection;
    }
        );

var app = builder.Build();

app.MapGet("/", async (GetConnection connectionGetter) =>
{
    using var con = await connectionGetter();
    return con.GetAll<Post>().ToList();
});

app.Run();

[Table("posts")]
public record Post(int Id, string Title, string Content);

public delegate Task<IDbConnection> GetConnection();
