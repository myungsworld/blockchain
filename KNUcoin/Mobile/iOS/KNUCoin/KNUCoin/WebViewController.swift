//
//  WebViewController.swift
//  KNUCoin
//
//  Created by Sanghyun Byun on 2020/05/06.
//  Copyright © 2020 sbyun. All rights reserved.
//

import UIKit

class WebViewController: UIViewController {
    
    @IBOutlet weak var webView: UIWebView!
    var url = URL(string: "")
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        let urlreq = URLRequest(url: url!)
        webView.loadRequest(urlreq)
    }
}
