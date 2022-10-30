import { Component, OnInit } from '@angular/core';
import { Character } from '../character';
import { CharacterService } from '../character.service';
import { Title } from '@angular/platform-browser';

@Component({
  selector: 'ig-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.css'],
})
export class MainComponent implements OnInit {
  title: string = '';

  selected: Character = { name: '' };

  histories: Character[] = [];

  url: string = encodeURIComponent(location.href);

  constructor(
    private sharedTitle: Title,
    private characterService: CharacterService
  ) {}

  ngOnInit(): void {
    this.title = this.sharedTitle.getTitle();
  }

  generate(): void {
    this.characterService.get().subscribe((c) => {
      this.histories.unshift(c);
      this.select(c);
    });
  }

  select(c: Character): void {
    this.selected = c;
  }
}
