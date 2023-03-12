package com.ellakaiser.dailynote.note;

import org.springframework.stereotype.Service;

import java.time.LocalDate;
import java.util.List;
@Service
public class InMemoryNoteService implements NoteService{
    private final InMemoryNoteDao dao;

    public InMemoryNoteService(InMemoryNoteDao dao) {
        this.dao = dao;
    }

    @Override
    public Note saveNote(Note note) {
        return this.dao.saveNote(note);
    }

    @Override
    public Note updateNote(Note note) {
        return this.dao.updateNote(note);
    }

    @Override
    public void deleteNote(String username, LocalDate date) {
        this.dao.deleteNote(username,date);
    }

    @Override
    public List<Note> findAllNotes() {
       return this.dao.findAllNotes();
    }

    @Override
    public Note findByDate(LocalDate date) {
        return this.dao.findByDate(date);
    }
}
