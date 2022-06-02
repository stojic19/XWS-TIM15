using AgentApplication.ClassLib.Service.Impl;
using Microsoft.EntityFrameworkCore.Migrations;

namespace AgentApplication.ClassLib.Migrations
{
    public partial class AdminSeed : Migration
    {
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            var salt1 = Encoder.CreateSalt(16);
            var password1 = Encoder.EncodePassword("Stojic", salt1);
            var salt2 = Encoder.CreateSalt(16);
            var password2 = Encoder.EncodePassword("Radisic", salt2);
            var salt3 = Encoder.CreateSalt(16);
            var password3 = Encoder.EncodePassword("Pesic", salt3);
            var salt4 = Encoder.CreateSalt(16);
            var password4 = Encoder.EncodePassword("Podunavac", salt4);
            migrationBuilder.Sql(
                @"INSERT INTO public."+ "\"Users\"" +
                "(\"Id\", \"Username\", \"Password\", \"PersonalInfo_FirstName\"," +
                "\"PersonalInfo_MiddleName\", \"PersonalInfo_LastName\", \"PersonalInfo_BirthDate\", \"PersonalInfo_Email\", " +
                "\"PersonalInfo_PhoneNumber\", \"TimeOfRegistration\", \"PersonalInfo_Gender\", \"Role\", \"Salt\") VALUES " +
                @"('815b90b9-7098-4cc6-b8bc-d25f373d418e', 'AleksaStojic','" + password1 + @"', 'Aleksa', 'Zeljko', 'Stojic', 
                '1999-1-19', 'stojic@gmail.com', '123456789', '2022-6-2 12:34:52', '0', '1', '" + salt1 + @"' ),
                ('cb120046-0ef7-4ddb-b04a-05ff2c2f49b6', 'AleksandarRadisic','" + password2 + @"', 'Aleksandar', 'Zoran', 'Radisic', 
                '1999-12-14', 'radisic@gmail.com', '123456789', '2022-6-2 12:34:52', '0', '1', '" + salt2 + @"' ),
                ('4492a33c-2f20-4177-86ab-692789c79666', 'AnjaPesic','" + password3 + @"', 'Anja', 'Milan', 'Pesic', 
                '1999-8-17', 'pesic@gmail.com', '123456789', '2022-6-2 12:34:52', '1', '1', '" + salt3 + @"' ),
                ('5bd7244b-bc44-4fa1-a2ce-f56a6339e21c', 'MilanPodunavac','" + password4 + @"', 'Milan', 'Nikola', 'Podunavac', 
                '2000-3-1', 'milancho@gmail.com', '123456789', '2022-6-2 12:34:52', '0', '1', '" + salt4 + @"' )"
            );
        }

        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.Sql("DELETE FROM public.\"Users\"");
        }
    }
}
