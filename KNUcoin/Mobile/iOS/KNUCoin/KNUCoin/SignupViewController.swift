//
//  SignupController.swift
//  KNUCoin
//
//  Created by Sanghyun Byun on 2020/04/30.
//  Copyright © 2020 sbyun. All rights reserved.
//

import UIKit

class SignupViewController: UIViewController {
    
    @IBOutlet weak var userid_textField: UITextField!
    @IBOutlet weak var userpwd_textField: UITextField!
    @IBOutlet weak var userRepwd_textField: UITextField!
    @IBOutlet weak var username_textField: UITextField!
    @IBOutlet weak var userPhone_textfield: UITextField!
    //인증번호
    
    var dupCheck: Bool = false
    var numCheck: Bool = false
    
    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view.
        
        //Titlebar configuration
        self.navigationItem.title = "회원가입"
    }

    override func touchesBegan(_ touches: Set<UITouch>, with event: UIEvent?){
          self.view.endEditing(true)
    }
    
    @IBAction func dupcheck_btn(_ sender: Any) {
        
        if let id = userid_textField.text {

            if(!id.isValid(String.ValidityType.id)){
                let alert = UIAlertController(title: "아이디를 확인하세요", message: "아이디를 확인하세요", preferredStyle: UIAlertController.Style.alert)
                let okAction = UIAlertAction(title: "Ok", style: .destructive)
                alert.addAction(okAction)
                self.present(alert, animated: true, completion: nil)

                return
            }
            
            let params = ["user_id" : id]

            guard let url = URL(string: "http://54.180.89.38/signup/check") else { return }
            var request = URLRequest(url: url)
            request.httpMethod = "POST"
            request.addValue("application/json", forHTTPHeaderField: "Content-Type")
            request.addValue("application/json", forHTTPHeaderField: "Accept")
            
            guard let httpBody = try? JSONSerialization.data(withJSONObject: params, options: []) else {return}
            request.httpBody = httpBody
            
            let task = URLSession.shared.dataTask(with: request) { (data, response, error) in
                
                if let e = error {
                    NSLog("An error has occured : \(e.localizedDescription)")
                    return
                }
                
                if let data = data {
                    DispatchQueue.main.async {
                        do {
                            let object = try JSONSerialization.jsonObject(with: data, options: []) as? NSDictionary
                            guard let result = object else {return}
                            guard let rs = result["result"] as? String else {return}
                            
                            if rs == "no duplication" {
                                
                                self.dupCheck = true
                                
                                let alert = UIAlertController(title: "사용 가능합니다!!", message: "사용 가능합니다!!", preferredStyle: UIAlertController.Style.alert)
                                let okAction = UIAlertAction(title: "Ok", style: .default)
                                alert.addAction(okAction)
                                self.present(alert, animated: true, completion: nil)
                                
                                
                            } else { //duplication
                                
                                self.dupCheck = false
                                
                                let alert = UIAlertController(title: "사용 불가능합니다!!", message: "사용 불가능합니다!!", preferredStyle: UIAlertController.Style.alert)
                                let okAction = UIAlertAction(title: "Ok", style: .destructive)
                                alert.addAction(okAction)
                                self.present(alert, animated: true, completion: nil)

                            }
                            
                        } catch {
                            print(error)
                        }
                    }
                }
            }
            
            task.resume()
        } else {return}
        
        
    }
    
    @IBAction func validateReq_btn(_ sender: Any) {
    }
    
    @IBAction func validateNum_btn(_ sender: Any) {
        self.numCheck = true
    }
    
    @IBAction func signup_btn(_ sender: Any) {
        
        if(!self.dupCheck){
            let alert = UIAlertController(title: "중복체크를 해주세요", message: "중복체크를 해주세요", preferredStyle: UIAlertController.Style.alert)
            let okAction = UIAlertAction(title: "Ok", style: .destructive)
            alert.addAction(okAction)
            self.present(alert, animated: true, completion: nil)

            return
        }
        
        if let id = userid_textField.text,
            let pwd = userpwd_textField.text,
            let repwd = userRepwd_textField.text,
            let name = username_textField.text,
            let phone = userPhone_textfield.text {
            
            if(!id.isValid(String.ValidityType.id)){
                let alert = UIAlertController(title: "아이디를 확인하세요", message: "아이디를 확인하세요", preferredStyle: UIAlertController.Style.alert)
                let okAction = UIAlertAction(title: "Ok", style: .destructive)
                alert.addAction(okAction)
                self.present(alert, animated: true, completion: nil)

                return
            }
            
            if(!pwd.isValid(String.ValidityType.password) && (pwd == repwd)){
                let alert = UIAlertController(title: "비밀번호를 확인하세요", message: "비밀번호를 확인하세요", preferredStyle: UIAlertController.Style.alert)
                let okAction = UIAlertAction(title: "Ok", style: .destructive)
                alert.addAction(okAction)
                self.present(alert, animated: true, completion: nil)

                return
            }
            
            if(!name.isValid(String.ValidityType.name)){
                let alert = UIAlertController(title: "이름을 확인하세요", message: "이름을 확인하세요", preferredStyle: UIAlertController.Style.alert)
                let okAction = UIAlertAction(title: "Ok", style: .destructive)
                alert.addAction(okAction)
                self.present(alert, animated: true, completion: nil)

                return
            }
            
            if(!phone.isValid(String.ValidityType.phone)){
                let alert = UIAlertController(title: "번호를 확인하세요", message: "번호를 확인하세요", preferredStyle: UIAlertController.Style.alert)
                let okAction = UIAlertAction(title: "Ok", style: .destructive)
                alert.addAction(okAction)
                self.present(alert, animated: true, completion: nil)

                return
            }
            
            let params = ["user_id" : id,
                        "pwd" : pwd,
                        "category" : "0",
                        "name" : name,
                        "phone" : phone]

            guard let url = URL(string: "http://54.180.89.38/signup/signup") else { return }
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
                    DispatchQueue.main.async {
                        do {
                            let object = try JSONSerialization.jsonObject(with: data, options: []) as? NSDictionary
                            guard let result = object else { return }
                            guard let rs = result["result"] as? String else { return }
                            
                            if rs == "Success" {

                                let alert = UIAlertController(title: "성공", message: "회원가입 완료!", preferredStyle: UIAlertController.Style.alert)

                                let okAction = UIAlertAction(title: "Ok", style: .default){ (action) in

                                    if let contoller = self.storyboard?.instantiateViewController(withIdentifier: "LoginView") {
                                        //move(push -> navi) controller
                                        self.navigationController?.pushViewController(contoller, animated: true)
                                    }
                                }

                                alert.addAction(okAction)
                                self.present(alert, animated: true, completion: nil)
                            }

                        } catch {
                            print(error)
                        }
                    }
                }
            }

            task.resume()
        }
    }
}
