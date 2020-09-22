//
//  SignupController.swift
//  KNUCoin
//
//  Created by Sanghyun Byun on 2020/05/02.
//  Copyright © 2020 sbyun. All rights reserved.
//

import Foundation

class SignupController {
    
    let user_id: String
    let user_pwd: String
    
    init(user_id: String, user_pwd: String) {
        self.user_id = user_id
        self.user_pwd = user_pwd
    }
    
    func excute() {
        
        let params = ["user_id" : self.user_id,
                      "user_pwd" : self.user_pwd]
        
        guard let url = URL(string: "http://54.180.87.226/user/sign_up") else { return }
        var request = URLRequest(url: url)
        request.httpMethod = "POST"
        request.addValue("application/json", forHTTPHeaderField: "Content-Type")
        request.addValue("application/json", forHTTPHeaderField: "Accept")
        
        guard let httpBody = try? JSONSerialization.data(withJSONObject: params, options: []) else { return }

        request.httpBody = httpBody
        
        let task = URLSession.shared.dataTask(with: request) { (data, response, error) in
            
            if let e = error {
                NSLog("An error has ooccured : \(e.localizedDescription)")
                return
            }
            
            if let data = data {
                do {
                    let object = try JSONSerialization.jsonObject(with: data, options: []) as? NSDictionary
                    guard let result = object else { return }
                    guard let rs = result["result"] as? String else { return }
                    
                } catch {
                    print(error)
                }
            }
        }
        
        task.resume()
        
    }
}



