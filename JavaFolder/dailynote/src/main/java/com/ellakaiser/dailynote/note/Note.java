package com.ellakaiser.dailynote.note;

import java.time.LocalDate;

public class Note {
    private Long id;
    private String note;
    private LocalDate dot;

    public Note(Long id, String note, LocalDate dot) {
        this.id = id;
        this.note = note;
        this.dot = dot;
    }

    public Note() {}

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getNote() {
        return note;
    }

    public void setNote(String note) {
        this.note = note;
    }

    public LocalDate getDot() {
        return dot;
    }

    public void setDot(LocalDate dot) {
        this.dot = dot;
    }
}
