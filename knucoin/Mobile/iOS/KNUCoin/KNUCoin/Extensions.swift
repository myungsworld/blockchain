//
//  Extensions.swift
//  KNUCoin
//
//  Created by Sanghyun Byun on 2020/05/11.
//  Copyright © 2020 sbyun. All rights reserved.
//

import Foundation

extension String {
    
    enum ValidityType {
        case id
        case password
        case name
        case phone
    }
    
    enum Regex: String {
        case id = "[a-zA-Z0-9]{6,25}"
        case password = "^(?=.*[a-z])(?=.*\\d)(?=.*[$@$!%*?&#])[A-Za-z\\d$@$!%*?&#]{6,25}"
        case name = "^[가-힣]+$"
        case phone = "^[0-9]{3}[0-9]{4}[0-9]{4}$"
    }
    
    func isValid(_ validityType: ValidityType) -> Bool {
        
        let format = "SELF MATCHES %@"
        var regex = ""
        
        switch validityType {
        case .id:
            regex = Regex.id.rawValue
        case .password:
            regex = Regex.password.rawValue
        case .name:
            regex = Regex.name.rawValue
        case .phone:
            regex = Regex.phone.rawValue
        }
        
        return NSPredicate(format: format, regex).evaluate(with: self)
    }
}
