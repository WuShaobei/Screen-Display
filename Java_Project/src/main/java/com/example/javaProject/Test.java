package com.example.javaProject;


import org.springframework.web.bind.annotation.*;

import java.util.*;

@RestController
@RequestMapping("/v1")
public class Test {
    @GetMapping("/test")
    public String helloSpring(){
        return "hello Spring";
    }
}
