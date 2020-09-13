//
//  LoginContoller.swift
//  KNUCoin
//
//  Created by Sanghyun Byun on 2020/04/30.
//  Copyright © 2020 sbyun. All rights reserved.
//

import UIKit

class LoginViewController: UIViewController {
    
    @IBOutlet weak var userid_textField: UITextField!
    @IBOutlet weak var userpwd_textField: UITextField!
    
    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view.
        //let cookies = HTTPCookie.cookies(withResponseHeaderFields: ["login" : "yes"], for: <#T##URL#>)
    }
    
    override func touchesBegan(_ touches: Set<UITouch>, with event: UIEvent?){
          self.view.endEditing(true)
    }
    
    @IBAction func click_SignupText(_ sender: Any) {
        
        if let contoller = self.storyboard?.instantiateViewController(withIdentifier: "SignupView") {
            //move(push -> navi) controller
            self.navigationController?.pushViewController(contoller, animated: true)
        }
    }
    
    @IBAction func login_btn(_ sender: Any) {
        /* Login Process */
        if let id = userid_textField.text,
            let pwd = userpwd_textField.text{
            
            if(!id.isValid(String.ValidityType.id)){
                let alert = UIAlertController(title: "아이디를 확인하세요", message: "아이디를 확인하세요", preferredStyle: UIAlertController.Style.alert)
                let okAction = UIAlertAction(title: "Ok", style: .destructive)
                alert.addAction(okAction)
                self.present(alert, animated: true, completion: nil)

                return
            }
            
            if(!pwd.isValid(String.ValidityType.password)){
                let alert = UIAlertController(title: "비밀번호를 확인하세요", message: "비밀번호를 확인하세요", preferredStyle: UIAlertController.Style.alert)
                let okAction = UIAlertAction(title: "Ok", style: .destructive)
                alert.addAction(okAction)
                self.present(alert, animated: true, completion: nil)

                return
            }
            
            let params = ["user_id" : id,
                              "pwd" : pwd]
                
                guard let url = URL(string: "http://54.180.89.38/login") else { return }
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
                                    
                                    let alert = UIAlertController(title: "성공", message: "로그인 완료!", preferredStyle: UIAlertController.Style.alert)
                                    
                                    let okAction = UIAlertAction(title: "Ok", style: .default){ (action) in
                                        
                                        if let contoller = self.storyboard?.instantiateViewController(withIdentifier: "TabBar") {
                                            //move(push -> navi) controller
                                            self.navigationController?.pushViewController(contoller, animated: true)
                                        }
                                    }
                                    
                                    alert.addAction(okAction)
                                    self.present(alert, animated: true, completion: nil)
                                    
                                } else {
                                    
                                    let alert = UIAlertController(title: "아이디 또는 비밀번호가 유효하지 않습니다", message: "아이디 또는 비밀번호가 유효하지 않습니다", preferredStyle: UIAlertController.Style.alert)
                                    let okAction = UIAlertAction(title: "Ok", style: .destructive)
                                    alert.addAction(okAction)
                                    self.present(alert, animated: true, completion: nil)
                                    
                                    return
                                }
                                
                            } catch {
                                print(error)
                            }
                        }
                    }
                }
            
                task.resume()

            
        } else {
            //if there is no
        }
    }
}
