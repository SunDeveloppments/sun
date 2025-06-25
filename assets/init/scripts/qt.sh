#!/bin/bash

handle_error() {
    echo "Erreur: $1"
    exit 1
}

if [ "$(ls -A)" ]; then
    echo "Le répertoire courant n'est pas vide. Continuer ? (y/n)"
    read answer
    if [ "$answer" != "y" ]; then
        echo "Operation annulée."
        exit 1
    fi
fi

mkdir -p src include tests || handle_error "Erreur lors de la création des répertoires."

cat > MyQtProject.pro <<EOL || handle_error "Erreur lors de la création du fichier MyQtProject.pro."
# Project created by qt.sh

QT       += core gui

greaterThan(QT_MAJOR_VERSION, 4): QT += widgets

CONFIG += c++11

# You can make your code fail to compile if it uses deprecated APIs.
# In order to do so, uncomment the following line.
#DEFINES += QT_DISABLE_DEPRECATED_BEFORE=0x060000    # disables all the APIs deprecated before Qt 6.0.0

SOURCES += src/main.cpp

HEADERS += include/mainwindow.h

FORMS +=

# Default rules for deployment.
qnx: target.path = /tmp/$${TARGET}/bin
else: unix:!android: target.path = /opt/$${TARGET}/bin
!isEmpty(target.path): INSTALLS += target

EOL

mkdir -p src || handle_error "Erreur lors de la création du répertoire src."
cat > src/main.cpp <<EOL || handle_error "Erreur lors de la création du fichier main.cpp."
#include <QApplication>
#include <QLabel>

int main(int argc, char *argv[])
{
    QApplication app(argc, argv);
    QLabel *label = new QLabel("Hello, World!");
    label->show();
    return app.exec();
}
EOL

mkdir -p include || handle_error "Erreur lors de la création du répertoire include."
cat > include/mainwindow.h <<EOL || handle_error "Erreur lors de la création du fichier mainwindow.h."
#ifndef MAINWINDOW_H
#define MAINWINDOW_H

#include <QMainWindow>

class MainWindow : public QMainWindow
{
    Q_OBJECT

public:
    MainWindow(QWidget *parent = nullptr);
    ~MainWindow();
};

#endif // MAINWINDOW_H
EOL

echo "Qt project structure O.K."
echo "Use this to compile and execute:"
echo "1. qmake MyQtProject.pro"
echo "2. make"
echo "3. ./MyQtProject"
