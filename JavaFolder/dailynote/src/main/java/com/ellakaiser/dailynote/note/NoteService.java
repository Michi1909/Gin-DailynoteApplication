package com.ellakaiser.dailynote.note;
import java.time.LocalDate;
import java.util.List;
public interface NoteService {
    Note saveNote(Note note);
    Note updateNote(Note note);
    public void deleteNote(String username, LocalDate date);
    public List<Note> findAllNotes();
    Note findByDate(LocalDate date);
}
