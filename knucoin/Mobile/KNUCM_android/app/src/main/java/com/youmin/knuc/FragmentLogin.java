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
import androidx.appcompat.app.AppCompatActivity;
import androidx.fragment.app.Fragment;

import java.util.HashMap;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;
import retrofit2.Retrofit;
import retrofit2.converter.gson.GsonConverterFactory;

public class FragmentLogin extends Fragment {

    EditText user_id;
    EditText user_pwd;
    Button btn_login, btn_gosignup;
    CallBackFragment callBackFragment;
    RetrofitInterface retrofitInterface;




    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {

        View view = inflater.inflate(R.layout.login_fragment, container, false);
        retrofitInterface = APIClient.getClient().create(RetrofitInterface.class);

        user_id = view.findViewById(R.id.user_id);
        user_pwd = view.findViewById(R.id.user_pwd);
        btn_login = view.findViewById(R.id.btn_login);
        btn_gosignup = view.findViewById(R.id.btn_gosignup);


        btn_gosignup.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                if (callBackFragment != null) {
                    callBackFragment.changeFragment();
                }

            }
        });

        btn_login.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {



                HashMap<String, String> map = new HashMap<>();
                map.put("user_id", user_id.getText().toString());
                map.put("pwd", user_pwd.getText().toString());

                retrofitInterface.do_login(map).enqueue(new Callback<LoginResult>(){
                    @Override
                    public void onResponse(Call<LoginResult> call, Response<LoginResult> response) {
                        LoginResult result = response.body();
                        Toast.makeText(getContext(), result.getResult(), Toast.LENGTH_LONG).show();

                        if (callBackFragment != null) {
                            callBackFragment.topaymentFragment();
                        }


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

    public void setCallBackFragment(CallBackFragment callBackFragment) {
        this.callBackFragment = callBackFragment;
    }

}
