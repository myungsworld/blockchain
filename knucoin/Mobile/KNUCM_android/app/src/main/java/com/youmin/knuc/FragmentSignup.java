package com.youmin.knuc;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Button;
import android.widget.EditText;
import android.widget.Toast;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.appcompat.app.AlertDialog;
import androidx.fragment.app.Fragment;

import com.youmin.knuc.POJO.User;

import java.util.HashMap;
import java.util.List;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;
import retrofit2.Retrofit;
import retrofit2.converter.gson.GsonConverterFactory;


public class FragmentSignup extends Fragment {


    EditText user_id, user_password, password_check, user_name, user_phone, edit_phonecheck;
    Button btn_duplicateId, btn_phonecheck, btn_confirm, btn_sign_up;
    RetrofitInterface retrofitInterface;


    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.signup_fragment, container, false);


        user_id = view.findViewById(R.id.user_id);
        user_password = view.findViewById(R.id.uesr_password);
        password_check = view.findViewById(R.id.password_check);
        user_name = view.findViewById(R.id.user_name);
        user_phone = view.findViewById(R.id.user_phone);
        edit_phonecheck = view.findViewById(R.id.edit_phonecheck);
        btn_duplicateId = view.findViewById(R.id.btn_duplicateId);
        btn_phonecheck = view.findViewById(R.id.btn_confirm);
        btn_confirm = view.findViewById(R.id.btn_confirm);
        btn_sign_up = view.findViewById(R.id.btn_signup);

        btn_sign_up.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {

                retrofitInterface = APIClient.getClient().create(RetrofitInterface.class);


                HashMap<String, String> map = new HashMap<>();
                map.put("user_id", user_id.getText().toString());
                map.put("pwd", user_password.getText().toString());
                map.put("user_category", "0");
                map.put("user_name", user_name.getText().toString());
                map.put("user_phone", user_phone.getText().toString());


                retrofitInterface.do_signup(map).enqueue(new Callback<LoginResult>() {
                    @Override
                    public void onResponse(Call<LoginResult> call, Response<LoginResult> response) {
                        LoginResult result = response.body();
                        Toast.makeText(getContext(), result.getResult(), Toast.LENGTH_LONG).show();
                        //통신 성공
                    }

                    @Override
                    public void onFailure(Call<LoginResult> call, Throwable t) {
                        //통신 실패
                        Toast.makeText(getContext(), t.getMessage(), Toast.LENGTH_LONG).show();
                    }

                });

            }
        });
        return view;
    }


}

