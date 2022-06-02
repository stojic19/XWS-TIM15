using System;
using Microsoft.EntityFrameworkCore.Migrations;

namespace AgentApplication.ClassLib.Migrations
{
    public partial class RemovedLinesAppDbContext : Migration
    {
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropForeignKey(
                name: "FK_Comment_Companies_CompanyId",
                table: "Comment");

            migrationBuilder.DropForeignKey(
                name: "FK_Grade_Companies_CompanyId",
                table: "Grade");

            migrationBuilder.DropForeignKey(
                name: "FK_JobOffer_Companies_CompanyId",
                table: "JobOffer");

            migrationBuilder.DropPrimaryKey(
                name: "PK_JobOffer",
                table: "JobOffer");

            migrationBuilder.DropPrimaryKey(
                name: "PK_Grade",
                table: "Grade");

            migrationBuilder.DropPrimaryKey(
                name: "PK_Comment",
                table: "Comment");

            migrationBuilder.AlterColumn<Guid>(
                name: "CompanyId",
                table: "JobOffer",
                type: "uuid",
                nullable: true,
                oldClrType: typeof(Guid),
                oldType: "uuid");

            migrationBuilder.AlterColumn<Guid>(
                name: "CompanyId",
                table: "Grade",
                type: "uuid",
                nullable: true,
                oldClrType: typeof(Guid),
                oldType: "uuid");

            migrationBuilder.AlterColumn<Guid>(
                name: "CompanyId",
                table: "Comment",
                type: "uuid",
                nullable: true,
                oldClrType: typeof(Guid),
                oldType: "uuid");

            migrationBuilder.AddPrimaryKey(
                name: "PK_JobOffer",
                table: "JobOffer",
                column: "Id");

            migrationBuilder.AddPrimaryKey(
                name: "PK_Grade",
                table: "Grade",
                column: "Id");

            migrationBuilder.AddPrimaryKey(
                name: "PK_Comment",
                table: "Comment",
                column: "Id");

            migrationBuilder.CreateIndex(
                name: "IX_JobOffer_CompanyId",
                table: "JobOffer",
                column: "CompanyId");

            migrationBuilder.CreateIndex(
                name: "IX_Grade_CompanyId",
                table: "Grade",
                column: "CompanyId");

            migrationBuilder.CreateIndex(
                name: "IX_Comment_CompanyId",
                table: "Comment",
                column: "CompanyId");

            migrationBuilder.AddForeignKey(
                name: "FK_Comment_Companies_CompanyId",
                table: "Comment",
                column: "CompanyId",
                principalTable: "Companies",
                principalColumn: "Id",
                onDelete: ReferentialAction.Restrict);

            migrationBuilder.AddForeignKey(
                name: "FK_Grade_Companies_CompanyId",
                table: "Grade",
                column: "CompanyId",
                principalTable: "Companies",
                principalColumn: "Id",
                onDelete: ReferentialAction.Restrict);

            migrationBuilder.AddForeignKey(
                name: "FK_JobOffer_Companies_CompanyId",
                table: "JobOffer",
                column: "CompanyId",
                principalTable: "Companies",
                principalColumn: "Id",
                onDelete: ReferentialAction.Restrict);
        }

        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropForeignKey(
                name: "FK_Comment_Companies_CompanyId",
                table: "Comment");

            migrationBuilder.DropForeignKey(
                name: "FK_Grade_Companies_CompanyId",
                table: "Grade");

            migrationBuilder.DropForeignKey(
                name: "FK_JobOffer_Companies_CompanyId",
                table: "JobOffer");

            migrationBuilder.DropPrimaryKey(
                name: "PK_JobOffer",
                table: "JobOffer");

            migrationBuilder.DropIndex(
                name: "IX_JobOffer_CompanyId",
                table: "JobOffer");

            migrationBuilder.DropPrimaryKey(
                name: "PK_Grade",
                table: "Grade");

            migrationBuilder.DropIndex(
                name: "IX_Grade_CompanyId",
                table: "Grade");

            migrationBuilder.DropPrimaryKey(
                name: "PK_Comment",
                table: "Comment");

            migrationBuilder.DropIndex(
                name: "IX_Comment_CompanyId",
                table: "Comment");

            migrationBuilder.AlterColumn<Guid>(
                name: "CompanyId",
                table: "JobOffer",
                type: "uuid",
                nullable: false,
                defaultValue: new Guid("00000000-0000-0000-0000-000000000000"),
                oldClrType: typeof(Guid),
                oldType: "uuid",
                oldNullable: true);

            migrationBuilder.AlterColumn<Guid>(
                name: "CompanyId",
                table: "Grade",
                type: "uuid",
                nullable: false,
                defaultValue: new Guid("00000000-0000-0000-0000-000000000000"),
                oldClrType: typeof(Guid),
                oldType: "uuid",
                oldNullable: true);

            migrationBuilder.AlterColumn<Guid>(
                name: "CompanyId",
                table: "Comment",
                type: "uuid",
                nullable: false,
                defaultValue: new Guid("00000000-0000-0000-0000-000000000000"),
                oldClrType: typeof(Guid),
                oldType: "uuid",
                oldNullable: true);

            migrationBuilder.AddPrimaryKey(
                name: "PK_JobOffer",
                table: "JobOffer",
                columns: new[] { "CompanyId", "Id" });

            migrationBuilder.AddPrimaryKey(
                name: "PK_Grade",
                table: "Grade",
                columns: new[] { "CompanyId", "Id" });

            migrationBuilder.AddPrimaryKey(
                name: "PK_Comment",
                table: "Comment",
                columns: new[] { "CompanyId", "Id" });

            migrationBuilder.AddForeignKey(
                name: "FK_Comment_Companies_CompanyId",
                table: "Comment",
                column: "CompanyId",
                principalTable: "Companies",
                principalColumn: "Id",
                onDelete: ReferentialAction.Cascade);

            migrationBuilder.AddForeignKey(
                name: "FK_Grade_Companies_CompanyId",
                table: "Grade",
                column: "CompanyId",
                principalTable: "Companies",
                principalColumn: "Id",
                onDelete: ReferentialAction.Cascade);

            migrationBuilder.AddForeignKey(
                name: "FK_JobOffer_Companies_CompanyId",
                table: "JobOffer",
                column: "CompanyId",
                principalTable: "Companies",
                principalColumn: "Id",
                onDelete: ReferentialAction.Cascade);
        }
    }
}
