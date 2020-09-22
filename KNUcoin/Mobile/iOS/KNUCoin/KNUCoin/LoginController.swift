//
//  LoginController.swift
//  KNUCoin
//
//  Created by Sanghyun Byun on 2020/05/01.
//  Copyright © 2020 sbyun. All rights reserved.
//

import Foundation

class LoginController {
    
    var user_id: String
    var user_pwd: String
    
    init(user_id: String, user_pwd: String) {
        self.user_id = user_id
        self.user_pwd = user_pwd
    }
    
    func show(){
        print(user_id, user_pwd)
    }
}
