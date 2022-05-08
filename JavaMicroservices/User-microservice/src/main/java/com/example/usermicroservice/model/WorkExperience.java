package com.example.usermicroservice.model;

import java.util.Date;

public class WorkExperience {

    private String companyName;
    private String jobTitle;
    private Date startDate;
    private Date endDate;

    public WorkExperience(String companyName, String jobTitle, Date startDate, Date endDate) {
        this.companyName = companyName;
        this.jobTitle = jobTitle;
        this.startDate = startDate;
        this.endDate = endDate;
    }

    public String getCompanyName() {
        return companyName;
    }

    public void setCompanyName(String companyName) {
        this.companyName = companyName;
    }

    public String getJobTitle() {
        return jobTitle;
    }

    public void setJobTitle(String jobTitle) {
        this.jobTitle = jobTitle;
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
