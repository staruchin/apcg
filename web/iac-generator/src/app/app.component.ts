import { Component, OnInit } from '@angular/core';
import { Title } from '@angular/platform-browser';

@Component({
  selector: 'ig-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
})
export class AppComponent implements OnInit {
  constructor(private sharedTitle: Title) {}

  ngOnInit(): void {
    this.sharedTitle.setTitle('ア〇パンマンっぽいキャラジェネレーター');
  }
}
