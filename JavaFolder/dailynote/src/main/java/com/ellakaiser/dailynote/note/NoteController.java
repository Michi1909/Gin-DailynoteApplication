package com.ellakaiser.dailynote.note;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/api/v1/notes")
public class NoteController {
    @GetMapping
    public List<String> findAllNotes(){
        return List.of("Ella","Hello World!");
    }
}