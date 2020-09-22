package com.youmin.knuc;

import android.app.Activity;
import android.os.Bundle;

public class FragmentSplash extends Activity {

    MainActivity mainActivity;


    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.splash_fragment);

        if (savedInstanceState == null) {
            mainActivity.addFragment();
        }
    }
}
