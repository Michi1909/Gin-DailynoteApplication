package com.ellakaiser.dailynote.note;

import java.time.LocalDate;

public class Note {
    private String username;
    private String notefromUser;
    private final LocalDate dot = LocalDate.now();;

    public Note(String username, String notefromUser) {
        this.username=username;
        this.notefromUser = notefromUser;
    }

    public Note() {}

    public String getNote() {
        return notefromUser;
    }

    public void setNote(String notefromUser) {
        this.notefromUser = notefromUser;
    }

    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public LocalDate getDot() {
        return dot;
    }

}
