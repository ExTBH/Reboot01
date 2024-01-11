#include <ncurses.h>
#include <menu.h>
#include <string.h>

#define MAX_FIELD_SIZE 20

typedef struct {
    char username[MAX_FIELD_SIZE + 1];
    char password[MAX_FIELD_SIZE + 1];
} Credentials;

void handle_resize(int signal) {
    // Handle resizing logic if needed
    // For simplicity, do nothing in this example
}

void initialize_ncurses() {
    initscr();
    cbreak();
    noecho();
    keypad(stdscr, TRUE);
    start_color(); // Enable color if not already enabled
    init_pair(1, COLOR_WHITE, COLOR_BLACK); // Define a color pair with white text on black background
    bkgd(COLOR_PAIR(1)); // Set the background color to black
}

void draw_login_page() {
    mvprintw(5, 5, "Username:");
    mvprintw(6, 5, "Password:");
    mvprintw(8, 5, "[Login]");
    mvprintw(8, 15, "[Register]");
}

void draw_register_page() {
    mvprintw(5, 5, "New Username:");
    mvprintw(6, 5, "New Password:");
    mvprintw(8, 5, "[Back]");
    mvprintw(8, 15, "[Create Account]");
}

void draw_input_fields(Credentials *credentials, int y, int x) {
    mvprintw(y, x, "%s", credentials->username);
    mvprintw(y + 1, x, "%s", credentials->password);
}

void handle_input(char *field, int max_size, int y, int x) {
    echo();
    curs_set(1); // Show cursor
    move(y, x);
    getnstr(field, max_size);
    curs_set(0); // Hide cursor
    noecho();
}

int main() {
    initialize_ncurses();

    // Set up the signal handler for SIGWINCH
    signal(SIGWINCH, handle_resize);

    char *choices[] = {"Login", "Register"};
    int num_choices = sizeof(choices) / sizeof(choices[0]);
    MENU *my_menu = create_menu(choices, num_choices);

    int c;
    Credentials credentials = {0}; // Initialize credentials struct with zeros

    while ((c = getch()) != KEY_F(1)) {
        switch (c) {
            case KEY_DOWN:
                menu_driver(my_menu, REQ_DOWN_ITEM);
                break;
            case KEY_UP:
                menu_driver(my_menu, REQ_UP_ITEM);
                break;
            case 10: // Enter key
                clear();
                if (item_index(current_item(my_menu)) == 0) {
                    draw_login_page();
                    draw_input_fields(&credentials, 5, 15);
                } else if (item_index(current_item(my_menu)) == 1) {
                    draw_register_page();
                    draw_input_fields(&credentials, 5, 20);
                }
                break;
        }
    }

    // Clean up
    
    ITEM **my_items = menu_items(my_menu);
    cleanup_menu(my_menu, my_items, num_choices);
    endwin();

    return 0;
}
