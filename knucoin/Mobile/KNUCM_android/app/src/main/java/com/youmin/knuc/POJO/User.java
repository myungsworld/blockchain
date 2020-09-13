package com.youmin.knuc.POJO;

public class User {

    private final String user_id;
    private final String user_pwd;
    private final String user_category;
    private final String user_name;
    private final String user_phone;


    public User(String user_id, String user_pwd, String user_category, String user_name, String user_phone) {
        this.user_id = user_id;
        this.user_pwd = user_pwd;
        this.user_category = user_category;
        this.user_name = user_name;
        this.user_phone = user_phone;
    }

    public String getUser_id() {
        return user_id;
    }

    public String getUser_pwd() {
        return user_pwd;
    }

    public String getUser_category() {
        return user_category;
    }

    public String getUser_name() {
        return user_name;
    }

    public String getUser_phone() {
        return user_phone;
    }

    private String result;

    public String getResult() {
        return result;
    }

    public void setResult(String result) {
        this.result = result;
    }
}
