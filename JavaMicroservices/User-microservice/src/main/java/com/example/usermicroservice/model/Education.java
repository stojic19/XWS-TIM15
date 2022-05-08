package com.example.usermicroservice.model;

import java.util.Date;

public class Education {

    private String institutionType;
    private String institutionName;
    private String title;
    private double gpa;
    private Date startDate;
    private Date endDate;

    public Education(String institutionType, String institutionName, String title, double gpa, Date startDate, Date endDate) {
        this.institutionType = institutionType;
        this.institutionName = institutionName;
        this.title = title;
        this.gpa = gpa;
        this.startDate = startDate;
        this.endDate = endDate;
    }

    public String getInstitutionType() {
        return institutionType;
    }

    public void setInstitutionType(String institutionType) {
        this.institutionType = institutionType;
    }

    public String getInstitutionName() {
        return institutionName;
    }

    public void setInstitutionName(String institutionName) {
        this.institutionName = institutionName;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public double getGpa() {
        return gpa;
    }

    public void setGpa(double gpa) {
        this.gpa = gpa;
    }

    public Date getStartDate() {
        return startDate;
    }

    public void setStartDate(Date startDate) {
        this.startDate = startDate;
    }

    public Date getEndDate() {
        return endDate;
    }

    public void setEndDate(Date endDate) {
        this.endDate = endDate;
    }
}
