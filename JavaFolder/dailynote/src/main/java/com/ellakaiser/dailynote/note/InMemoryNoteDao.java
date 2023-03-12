package com.ellakaiser.dailynote.note;

import org.springframework.stereotype.Repository;

import java.time.LocalDate;
import java.util.ArrayList;
import java.util.List;
import java.util.stream.IntStream;
@Repository
public class InMemoryNoteDao {
    private final List<Note> NOTES = new ArrayList<>();

    public Note saveNote(Note note) {
        if(!NOTES.contains(note.getUsername())){
        NOTES.add(note);
            return note;
        }
        return null;
    }

    public Note updateNote(Note note) {
        var noteIndex = IntStream.range(0,NOTES.size()).filter(index -> NOTES.get(index).equals(note)).findFirst().orElse(-1);
        if(noteIndex>-1){

            if(NOTES.get(noteIndex).getDot().equals(LocalDate.now())){
            NOTES.set(noteIndex,note);
            return note;
            }

            return null;
        }
        return null;
    }

    public List<Note> findAllNotes() {
        return NOTES;
    }

    public Note findByDate(LocalDate date) {
        return NOTES.stream().filter(note -> date.equals(note.getDot())).findAny().orElse(null);
    }

    public void deleteNote(String username, LocalDate date) {
        var noteUser = NOTES.stream().filter(note -> username.equals(note.getUsername()));
        noteUser.filter(note -> date.equals(note.getDot()));;
        NOTES.remove(noteUser);
    }
}
