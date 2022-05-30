using Microsoft.EntityFrameworkCore.Migrations;

namespace AgentApplication.ClassLib.Migrations
{
    public partial class CompanyInfoUpdate : Migration
    {
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.RenameColumn(
                name: "PhoneNumber",
                table: "Companies",
                newName: "CompanyInfo_PhoneNumber");

            migrationBuilder.RenameColumn(
                name: "Name",
                table: "Companies",
                newName: "CompanyInfo_Name");

            migrationBuilder.RenameColumn(
                name: "Email",
                table: "Companies",
                newName: "CompanyInfo_Email");

            migrationBuilder.RenameColumn(
                name: "Description",
                table: "Companies",
                newName: "CompanyInfo_Description");

            migrationBuilder.RenameColumn(
                name: "Culture",
                table: "Companies",
                newName: "CompanyInfo_Culture");

            migrationBuilder.RenameColumn(
                name: "Address",
                table: "Companies",
                newName: "CompanyInfo_Address");
        }

        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.RenameColumn(
                name: "CompanyInfo_PhoneNumber",
                table: "Companies",
                newName: "PhoneNumber");

            migrationBuilder.RenameColumn(
                name: "CompanyInfo_Name",
                table: "Companies",
                newName: "Name");

            migrationBuilder.RenameColumn(
                name: "CompanyInfo_Email",
                table: "Companies",
                newName: "Email");

            migrationBuilder.RenameColumn(
                name: "CompanyInfo_Description",
                table: "Companies",
                newName: "Description");

            migrationBuilder.RenameColumn(
                name: "CompanyInfo_Culture",
                table: "Companies",
                newName: "Culture");

            migrationBuilder.RenameColumn(
                name: "CompanyInfo_Address",
                table: "Companies",
                newName: "Address");
        }
    }
}
