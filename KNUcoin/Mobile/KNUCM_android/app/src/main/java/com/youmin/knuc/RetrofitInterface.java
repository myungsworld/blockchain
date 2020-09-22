package com.youmin.knuc;

import android.util.Log;
import android.widget.EditText;

import com.youmin.knuc.POJO.User;

import java.util.HashMap;

import retrofit2.Call;
import retrofit2.http.Body;
import retrofit2.http.Field;
import retrofit2.http.FieldMap;
import retrofit2.http.FormUrlEncoded;
import retrofit2.http.GET;
import retrofit2.http.POST;

public interface RetrofitInterface {


    @POST("/api/user/login")
    Call<LoginResult> do_login(@Body HashMap<String, String> map);


   // @FormUrlEncoded
    @POST("/api/user/signup")
    Call<LoginResult> do_signup(@Body HashMap<String, String> map);

   // @POST("/check")
  //  Call<>

    //@GET("api/user/autologin")
    //Call<Void> auto_login()



}
