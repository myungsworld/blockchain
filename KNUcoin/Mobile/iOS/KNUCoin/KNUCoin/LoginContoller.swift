//
//  LoginContoller.swift
//  KNUCoin
//
//  Created by Sanghyun Byun on 2020/04/30.
//  Copyright © 2020 sbyun. All rights reserved.
//

import UIKit

class LoginViewController: UIViewController {
    
    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view.
    }
    
    @IBAction func click_SignupText(_ sender: Any) {
        
        
        if let contoller = self.storyboard?.instantiateViewController(withIdentifier: "SignupController") {
            //move(push -> navi) controller
            self.navigationController?.pushViewController(contoller, animated: true)
        }
    }
    
}
