package com.ellakaiser.dailynote.note;

import org.springframework.stereotype.Service;

import java.time.LocalDate;
import java.util.List;

@Service
public class NoteService {

    public List<Note> findAllNotes(){
        return List.of(new Note(1L,"This is the first Note", LocalDate.now()),
                new Note(2L,"This is the secound Note", LocalDate.now()));
    }
}
