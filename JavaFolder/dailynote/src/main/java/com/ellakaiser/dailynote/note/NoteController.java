package com.ellakaiser.dailynote.note;

import org.springframework.web.bind.annotation.*;

import java.time.LocalDate;
import java.util.List;

@RestController
@RequestMapping("/api/v1/notes")
public class NoteController {

    private NoteService noteService;

    public NoteController(NoteService noteService) {
        this.noteService = noteService;
    }
    @PostMapping
    public Note saveNote(@RequestBody Note note) {
        return noteService.saveNote(note);
    }
    @PutMapping
    public Note updateNote(@RequestBody Note note) {
        return noteService.updateNote(note);
    }
    @GetMapping
    public List<Note> findAllNotes() {
        return noteService.findAllNotes();
    }
    @GetMapping("/{date}")
    public Note findByDate(@PathVariable("date") LocalDate date) {
        return noteService.findByDate(date);
    }
    @DeleteMapping("/{username}/{date}")
    public void deleteNote(@PathVariable String username, @PathVariable LocalDate date) {
       noteService.deleteNote(username,date);
    }

}
